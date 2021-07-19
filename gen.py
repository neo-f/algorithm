from dataclasses import dataclass
import os
from typing import List
import json
import datetime

GROUPS = dict(
    week1="88,21,26,283,66,206,20,155,150,641,141,227,560,25,1248,304,1109,53,1,167,15,11,84,239,42,560",
    week2="874,49,30,146,78,77,46,111,104,226,98,50,22",
    week3="106,210,130,94,589,429,297,105,1245,236,684,207,17,51,200,433,329",
    week4="23,295,355,239,704,410,162,34",
    week5="912,1122",
)

TITLE = "# Leetcode Solutions\n\n"


with open("qs.json", "r") as f:
    DB = json.load(f)


@dataclass
class Item:
    title: str
    slug: str
    tags: List[str]
    diffcult: str
    solutions_dir: str


STATISTIC = {"solved": 0, "lang": {}}


def main():
    content = ""
    for week, nums in GROUPS.items():
        items = _parse_week(nums)
        content += _gen_week(week, items)

    with open("README.md", "w") as readme:
        readme.write(TITLE)
        # 统计信息
        readme.write(_gen_statistic())
        readme.write(content)


def _parse_week(nums: str):
    items = []
    for num in nums.split(","):
        qid = num.zfill(4)
        q = DB[qid]
        folder = f"solutions/{qid}-{q['title']}"
        if not os.path.isdir(folder):
            os.mkdir(folder)
        items.append(
            Item(
                title=f"{qid}-{q['title']}",
                slug=q["slug"],
                tags=q["tags"],
                diffcult=q["difficulty"],
                solutions_dir=folder,
            )
        )
    return items


EXT = {
    "go": "Go",
    "py": "Python",
    "rs": "Rust",
    "js": "Javascript",
}
DIFFICULTY = {
    "EASY": "🟩",
    "MEDIUM": "🟨",
    "HARD": "🟥",
}


def _gen_statistic() -> str:
    content = ""
    content += f"> Generate Time: {datetime.date.today()}\t(Total Solved: {STATISTIC['solved']})\n"
    for lang, num in STATISTIC["lang"].items():
        content += f"> + {lang} Solutions: {num}\n"
    return content


def _gen_week(week: str, items: List[Item]) -> str:
    ret = "***\n"
    ret += f"## {week.capitalize()}\n\n"
    titles = ["No.", "题目", "难度", "Tags", "Solutions"]
    ret += f"|{'|'.join(titles)}|\n"
    ret += "|---|---|---|---|---|\n"
    for idx, item in enumerate(items, start=1):
        line = []
        line.append(str(idx))
        line.append(f"[{item.title}](https://leetcode-cn.com/problems/{item.slug})")
        line.append(f"{DIFFICULTY[item.diffcult]}")
        line.append(" ".join([f"`{t}`" for t in item.tags]))
        line.append(_gen_solutions_str(item.solutions_dir))
        ret += f"|{'|'.join(line)}|\n"
    ret += "\n\n"
    return ret


def _gen_solutions_str(solutions_dir: str) -> str:
    ret = []
    count = {}
    solutions = os.listdir(solutions_dir)
    STATISTIC["solved"] += 1
    for s in solutions:
        lang = EXT[s.split(".")[-1]]
        count.setdefault(lang, [])
        count[lang].append(s)

    for lang, solutions in count.items():
        STATISTIC["lang"].setdefault(lang, 0)
        STATISTIC["lang"][lang] += len(solutions)
        if len(count[lang]) == 1:
            href = f"{solutions_dir}/{solutions[0]}".replace(" ", "%20")
            ret.append(f"[{lang}]({href})")
        else:
            for idx, solution in enumerate(solutions, start=1):
                href = f"{solutions_dir}/{solution}".replace(" ", "%20")
                ret.append(f"[{lang}-{idx}]({href})")

    return " ".join(ret)


if __name__ == "__main__":
    main()