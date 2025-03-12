
const baseUrl = "http://localhost:8080/api/store"

export const buyCrate = async (jwt, name, count) => {
    const payload = {
        name: name,
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
