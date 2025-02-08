import { BrowserRouter as Router, Route, Routes } from "react-router"
import Navbar from "./components/Navbar"
import LandingPage from "./pages/LandingPage"
import SignUpLogin from "./pages/SignUpLogin"
import { useEffect, useState } from "react"
import useAuth from "./stores/authStore"
import Inventory from "./pages/Inventory"

function App() {
    const { loggedIn, setLoggedIn } = useAuth()
    const [loading, setLoading] = useState(true)

    const loadUser = async () => {
        // check if user is logged in
        const jwt = localStorage.getItem("jwt")
        if (jwt) {
            console.log("jwt exists")
            setLoggedIn(true)
        }

        setLoading(false)
    }

    useEffect(() => {
        loadUser()
    }, [])

    if (loading) return null

    return (
        <Router>
            <Navbar />
            <Routes>
                <Route
                    path="/"
                    element={<LandingPage />}
                />
                <Route
                    path="/login"
                    element={loggedIn ? <LandingPage /> : <SignUpLogin />}
                />
                <Route
                    path="/inventory"
                    element={loggedIn ? <Inventory /> : <SignUpLogin />}
                />
            </Routes>
        </Router>
    )
}

export default App
