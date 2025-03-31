
const baseUrl = "http://localhost:8080/v1/users/"

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
        const data = await res.json()
        return data
    } catch (error) {
        console.error("Error: ", error)
    }
}

export const getStats = async (jwt, userId) => {
    const opts = {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": "Bearer " + jwt,
        },
    }

    const url = baseUrl + userId + "/stats"

    try {
        const res = await fetch(url, opts)
        const data = await res.json()
        return data
    } catch (error) {
        throw error
    }
}

export const getRecentTradeups = async (jwt, userId) => {
    const opts = {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer " + jwt
        },
    }

    const url = baseUrl + userId + "/recent"

    try {
        const res = await fetch(url, opts)
        const data = await res.json()
        return data
    } catch (error) {
      console.error("Failed to fetch recent tradeups: ", error);
    }
}
