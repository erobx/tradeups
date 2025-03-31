import { useEffect, useRef, useState } from "react"
import { BrowserRouter as Router, Route, Routes } from "react-router"

import Navbar from "./components/Navbar"
import LandingPage from "./pages/LandingPage"
import SignUpLogin from "./pages/SignUpLogin"
import DashboardPage from "./pages/DashboardPage"
import ActiveTradeups from "./pages/ActiveTradeups"
import Tradeup from "./components/Tradeup/Tradeup"
import Store from "./pages/Store"
import Settings from "./pages/Settings"

import useAuth from "./stores/authStore"
import useUser from "./stores/userStore"
import useInventory from "./stores/inventoryStore"

import { getInventory } from "./api/inventory"
import { getUser } from "./api/user"
import { themeChange } from "theme-change"

function App() {
  const { loggedIn, setLoggedIn } = useAuth()
  const { user, setUser, setBalance } = useUser()
  const { inventory, setInventory, addItem, removeItem } = useInventory()
  const [loading, setLoading] = useState(true)

  const loadUser = async () => {
    // check if user is logged in
    const jwt = localStorage.getItem("jwt")
    if (jwt) {
      setLoggedIn(true)
      const userData = await getUser(jwt)
      setUser(userData.user)
      setBalance(userData.user.balance)
      await loadItems(jwt, userData.user.id)
    }
    setLoading(false)
  }

  const loadItems = async (jwt, userId) => {
    try {
      const data = await getInventory(jwt, userId)

      if (data.skins === "empty") {
        setInventory([])
      } else {
        data.skins = data.skins.map(skin => ({
          ...skin,
          price: parseFloat(skin.price).toFixed(2)
        }))
        setInventory(data.skins)
      }
    } catch (error) {
      console.error(error)
    }
  }

  useEffect(() => {
    loadUser()
    themeChange(true)
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
        <Route path="/settings" element={loggedIn ? <Settings /> : <SignUpLogin />} />
      </Routes>
    </Router>
  )
}

export default App
