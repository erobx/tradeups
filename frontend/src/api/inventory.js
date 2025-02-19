/*
 * API functions for a user's inventory
 *
 */

const baseUrl = "http://localhost:8080/api"

// api/users/:id/inventory
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
        const data = res.json()
        return data
    } catch (error) {
        console.error("Error: ", error)
    }
}
