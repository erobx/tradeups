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

export const getTradeup = async (id) => {
    const url = baseUrl + "/tradeups/" + id

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

export const addSkinToTradeup = async (jwt, invId, tradeupId) => {
    const url = baseUrl + "/tradeups/add"

    const headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer " + jwt,
    }

    const payload = {
        invId: invId,
        tradeupId: tradeupId,
    }

    try {
        const res = await fetch(url, {
            method: "PUT",
            headers: headers,
            body: JSON.stringify(payload),
        })
        return res
    } catch (error) {
        console.error("Error: ", error)
    }
}
