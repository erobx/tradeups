
const baseUrl = "http://localhost:8080/v1/store"

export const buyCrate = async (jwt, name, rarity, count) => {
    const payload = {
        name: name,
        rarity: rarity,
        count: count,
    }

    const opts = {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": "Bearer " + jwt,
        },
        body: JSON.stringify(payload),
    }

    try {
        const res = await fetch(baseUrl+"/buy", opts)
        const data = await res.json()
        return data
    } catch (error) {
        console.error(error)
    }
}
