import { useEffect, useState } from "react"
import SignUp from "../components/SignUp"
import Login from "../components/Login"

function SignUpLogin() {
    const [token, setToken] = useState('')

    useEffect(() => {
        const token = localStorage.getItem('jwt-token')
        setToken(token)
    }, [])

    if (!token) {
        return (
            <SignUp />
        )
    }

    return (
        <Login />
    )
}

export default SignUpLogin
