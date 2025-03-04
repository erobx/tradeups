/*
 *
 *
 */

const baseUrl = "http://localhost:8080/api"

export const getTradeups = async () => {
    const url = baseUrl + "/tradeups"

    const headers = {
        "Content-Type": "application/json",
    }

    try {
        const res = await fetch(url, {
            method: "GET",
            headers: headers,
        })
        const data = res.json()
        return data
    } catch (error) {
        console.error("Error: ", error)
    }
}
