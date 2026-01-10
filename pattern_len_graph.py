import pandas as pd
import matplotlib.pyplot as plt
from pathlib import Path


df = pd.read_csv("output/pattern_length.csv")

HASH_ORDER = ["sum", "poly_mod", "djb2", "poly_nomod"]

Path("plots").mkdir(exist_ok=True)
Path("plots/pattern").mkdir(exist_ok=True)

def plot_metric(text_name, y_column, ylabel, filename):
    subset = df[df["text_name"] == text_name]

    plt.figure(figsize=(15, 12))

    for hash_name in HASH_ORDER:
        data = subset[subset["hash_name"] == hash_name]
        if data.empty:
            continue

        plt.plot(
            data["pattern_length"],
            data[y_column],
            # marker="o",
            label=hash_name
        )

    plt.xlabel("Pattern length")
    plt.ylabel(ylabel)
    plt.title(f"{ylabel} vs pattern length ({text_name})")
    plt.legend()
    plt.grid(True)

    plt.tight_layout()
    plt.savefig(filename)
    plt.close()

for text in ["lorem", "alice", "dna", "repetitive_100k"]:
    plot_metric(
        text_name=text,
        y_column="time_ns",
        ylabel="Time (ns)",
        filename=f"plots/pattern/time_{text}.png"
    )

for text in ["lorem", "alice", "dna", "repetitive_100k"]:
    plot_metric(
        text_name=text,
        y_column="collisions",
        ylabel="Collisions",
        filename=f"plots/pattern/collisions_{text}.png"
    )

for text in ["lorem", "alice", "dna", "repetitive_100k"]:
    plot_metric(
        text_name=text,
        y_column="char_comparisons",
        ylabel="Сhar Сmparisons",
        filename=f"plots/pattern/char_comparisons_{text}.png"
    )


for text in ["lorem", "alice", "dna", "repetitive_100k"]:
    plot_metric(
        text_name=text,
        y_column="matches",
        ylabel="Matches",
        filename=f"plots/pattern/matches_{text}.png"
    )

for text in ["lorem", "alice", "dna", "repetitive_100k"]:
    plot_metric(
        text_name=text,
        y_column="matches",
        ylabel="Matches",
        filename=f"plots/pattern/matches_{text}.png"
    )

for text in ["lorem", "alice", "dna", "repetitive_100k"]:
    plot_metric(
        text_name=text,
        y_column="hash_matches",
        ylabel="Hash Matches",
        filename=f"plots/pattern/hash_matches_{text}.png"
    )

print("ok")
