import { useEffect, useState } from "react"
import { BrowserRouter as Router, Route, Routes } from "react-router"
import Navbar from "./components/Navbar"
import LandingPage from "./pages/LandingPage"
import SignUpLogin from "./pages/SignUpLogin"
import DashboardPage from "./pages/DashboardPage"
import ActiveTradeups from "./pages/ActiveTradeups"
import Tradeup from "./components/Tradeup"
import Store from "./pages/Store"
import useAuth from "./stores/authStore"
import useUserId from "./stores/userStore"
import useInventory from "./stores/inventoryStore"
import { getInventory } from "./api/inventory"

function App() {
  const { loggedIn, setLoggedIn } = useAuth()
  const { userId, setUserId } = useUserId()
  const { inventory, setInventory, addItem, removeItem } = useInventory()
  const [loading, setLoading] = useState(true)

  const loadUser = async () => {
    // check if user is logged in
    const jwt = localStorage.getItem("jwt")
    if (jwt) {
      console.log("jwt exists")
      setLoggedIn(true)
      const userId = localStorage.getItem("userId")
      if (userId) {
        setUserId(userId)
        loadItems(userId)
      }
    }

    setLoading(false)
  }

  const loadItems = async (userId) => {
    const jwt = localStorage.getItem("jwt")

    try {
      const data = await getInventory(jwt, userId)

      if (data.skins === "empty") {
        setInventory([])
      } else {
        setInventory(data.skins)
      }
    } catch (error) {
      console.error(error)
    }
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
        <Route path="/login/*" element={loggedIn ? <DashboardPage /> : <SignUpLogin />} />
        <Route path="/store" element={<Store />} />
        <Route path="/tradeups" element={<ActiveTradeups />} />
        <Route path="/tradeups/:tradeupId" element={<Tradeup />} />
        <Route path="/dashboard/*" element={loggedIn ? <DashboardPage /> : <SignUpLogin />} />
      </Routes>
    </Router>
  )
}

export default App
