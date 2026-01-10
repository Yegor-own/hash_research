import pandas as pd
import matplotlib.pyplot as plt
from pathlib import Path

df = pd.read_csv("output/text_length.csv")

HASH_ORDER = ["sum", "poly_mod", "djb2", "poly_nomod"]

Path("plots").mkdir(exist_ok=True)
Path("plots/text").mkdir(exist_ok=True)

random_df = df[df["text_name"].str.startswith("random")]

def plot_metric(y_column, ylabel, filename):
    plt.figure(figsize=(15, 12))

    for hash_name in HASH_ORDER:
        data = random_df[random_df["hash_name"] == hash_name]
        if data.empty:
            continue

        # сортировка по X
        data = data.sort_values("text_length")

        plt.plot(
            data["text_length"],
            data[y_column],
            # marker="o",
            label=hash_name
        )

    plt.xlabel("Text length")
    plt.ylabel(ylabel)
    plt.title(f"{ylabel} vs text length (random text)")
    plt.legend()
    plt.grid(True)

    plt.tight_layout()
    plt.savefig(filename)
    plt.close()


plot_metric(
    y_column="time_ns",
    ylabel="Time (ns)",
    filename="plots/text/time_vs_text_length.png"
)

plot_metric(
    y_column="collisions",
    ylabel="Collisions",
    filename="plots/text/collisions_vs_text_length.png"
)

plot_metric(
    y_column="char_comparisons",
    ylabel="char_comparisons",
    filename="plots/text/char_comparisons_vs_text_length.png"
)

plot_metric(
    y_column="hash_matches",
    ylabel="hash_matches",
    filename="plots/text/hash_matches_vs_text_length.png"
)

plot_metric(
    y_column="matches",
    ylabel="matches",
    filename="plots/text/matches_vs_text_length.png"
)

print("ok")
