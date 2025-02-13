export const wears = [
    "Battle-Scarred",
    "Well-Worn",
    "Field-Tested",
    "Minimal Wear",
    "Factory New",
]

export const wearsToAbbr = {
    "Battle-Scarred": "BS",
    "Well-Worn": "WW",
    "Field-Tested": "FT",
    "Minimal Wear": "MW",
    "Factory New": "FN",
}

// min inclusive, max exclusive
export const wearsToFloat = {
    "Battle-Scarred": {
        min: 0.45,
        max: 1.00,
    },
    "Well-Worn": {
        min: 0.37,
        max: 0.45,
    },
    "Field-Tested": {
        min: 0.15,
        max: 0.37,
    },
    "Minimal Wear": {
        min: 0.07,
        max: 0.15,
    },
    "Factory New": {
        min: 0.00,
        max: 0.07,
    },
}
