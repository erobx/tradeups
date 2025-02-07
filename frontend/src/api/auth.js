/*
 * API Functions for Auth
 *
 */

// {"username":"","email":"","password":""}
const submitSignup = async (username, email, password) => {
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
        const res = await fetch("http://localhost:8080/auth/register", opts)
        const data = await res.json()
        return data // { jwt: string }
    } catch (error) {
        console.error('Error:', error)
    }
}

export default submitSignup
