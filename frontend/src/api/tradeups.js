/*
 *
 *
 */

const baseUrl = "http://localhost:8080/api/tradeups"

export const getTradeups = async () => {
    const headers = {
        "Content-Type": "application/json",
    }

    try {
        const res = await fetch(baseUrl, {
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
    const url = baseUrl + "/" + id

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
    const url = baseUrl + "/add"

    const headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer " + jwt,
    }

    const payload = {
        invId: invId,
        tradeupId: tradeupId,
    }

    const opts = {
        method: "PUT",
        headers: headers,
        body: JSON.stringify(payload),
    }

    try {
        const res = await fetch(url, opts)
        return res
    } catch (error) {
        console.error("Error: ", error)
    }
}
