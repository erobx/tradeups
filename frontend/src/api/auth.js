/*
 * API functions for authentication
 *
 */

const baseUrl = "http://localhost:8080/auth"

// {"username":"","email":"","password":""}
export const submitSignup = async (username, email, password) => {
    const user = {
        username: username,
        email: email,
        password: password,
    }

    const opts = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(user)
    }

    try {
        const res = await fetch(baseUrl+"/register", opts)
        const data = await res.json()
        return data
    } catch (error) {
        console.error('Error:', error)
    }
}

export const submitLogin = async (email, password) => {
    const creds = {
        email: email,
        password: password,
    }

    const opts = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(creds)
    }

    try {
        const res = await fetch(baseUrl+"/login", opts)
        const data = await res.json()
        return data
    } catch (error) {
        console.error('Error:', error)
    }
}
