import { BrowserRouter as Router, Route, Routes } from "react-router"
import Navbar from "./components/Navbar"
import LandingPage from "./pages/LandingPage"
import SignUpLogin from "./pages/SignUpLogin"
import DashboardPage from "./pages/DashboardPage"
import TradeupsPage from "./pages/TradeupsPage"
import Tradeup from "./components/Tradeup"
import { useEffect, useState } from "react"
import useAuth from "./stores/authStore"

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
        <Route index element={<LandingPage />} />
        <Route path="/login" element={loggedIn ? <DashboardPage /> : <SignUpLogin />} />
        <Route path="/tradeups" element={<TradeupsPage />} />
        <Route path="/tradeups/:tradeupId" element={<Tradeup />} />
        <Route path="/dashboard/*" element={loggedIn ? <DashboardPage /> : <SignUpLogin />} />
      </Routes>
    </Router>
  )
}

export default App
