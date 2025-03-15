/*
 * API functions for a user's inventory
 *
 */

const baseUrl = "http://localhost:8080/api"

// api/users/:userId/inventory
export const getInventory = async (jwt, userId) => {
    const url = baseUrl + "/users/" + userId + "/inventory"

    const headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer " + jwt,
    }

    try {
        const res = await fetch(url, {
            method: "GET",
            headers: headers,
        })
        const data = await res.json()
        return data
    } catch (error) {
        console.error("Error: ", error)
    }
}

// api/users/:userId/inventory/:invId
export const deleteSkin = async (jwt, userId, invId) => {
    const url = baseUrl + "/users/" + userId + "/inventory/" + invId

    const headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer " + jwt,
    }

    try {
        const res = await fetch(url, {
            method: "DELETE",
            headers: headers,
        })
        return res
    } catch (error) {
        console.error("Error: ", error)
    }
}
