
const baseUrl = "http://localhost:8080/api/users"

export const getUser = async (jwt) => {
    const opts = {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": "Bearer " + jwt,
        }
    }

    try {
        const res = await fetch(baseUrl, opts)
        const data = res.json()
        return data
    } catch (error) {
        console.error("Error: ", error)
    }
}
