import { useEffect, useState } from "react"
import { BrowserRouter as Router, Route, Routes } from "react-router"
import Navbar from "./components/Navbar"
import LandingPage from "./pages/LandingPage"
import SignUpLogin from "./pages/SignUpLogin"
import DashboardPage from "./pages/DashboardPage"
import ActiveTradeups from "./pages/ActiveTradeups"
import Tradeup from "./components/Tradeup/Tradeup"
import Store from "./pages/Store"
import useAuth from "./stores/authStore"
import useUser from "./stores/userStore"
import useInventory from "./stores/inventoryStore"
import { getInventory } from "./api/inventory"
import { getUser } from "./api/user"

function App() {
  const { loggedIn, setLoggedIn } = useAuth()
  const { user, setUser, setBalance } = useUser()
  const { inventory, setInventory, addItem, removeItem } = useInventory()
  const [loading, setLoading] = useState(true)

  const loadUser = async () => {
    // check if user is logged in
    const jwt = localStorage.getItem("jwt")
    if (jwt) {
      console.log("jwt exists")
      setLoggedIn(true)
      const userData = await getUser(jwt)
      setUser(userData.user)
      loadItems(jwt, userData.user.id)
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
