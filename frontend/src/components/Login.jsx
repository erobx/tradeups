import { useState } from "react"
import { submitLogin } from "../api/auth"
import { useNavigate } from "react-router"
import useAuth from "../stores/authStore"
import useUserId from "../stores/userStore"
import { getInventory } from "../api/inventory"

function Login() {
  const navigate = useNavigate()
  const { loggedIn, setLoggedIn } = useAuth()
  const { userId, setUserId } = useUserId()
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [loading, setLoading] = useState(false)

  const loadItems = async (userId) => {
    const jwt = localStorage.getItem("jwt")

    try {
      const data = await getInventory(jwt, userId)
      data.skins.forEach(skin => {
        skin.skinPrice = parseFloat(skin.skinPrice).toFixed(2)
      })
      setInventory(data.skins)
    } catch (error) {
        console.error(error)
    }
  }

  const handleSubmit = async (e) => {
      if (loading) return
      e.preventDefault()
      setLoading(true)

      try {
        const data = await submitLogin(email, password)
        if (data) {
          setLoggedIn(true)
          localStorage.setItem("userId", data.userId)
          setUserId(data.userId)
          loadItems(data.userId)
          navigate("/dashboard")
          resetForm()
        } else {
          console.error("Login failed. Please check your credentials.")
        }
      } catch (error) {
          console.error("Error during login:", error)
      } finally {
          setLoading(false)
      }
  }

  const resetForm = () => {
    setEmail("")
    setPassword("")
  }

  return (
    <fieldset className="fieldset w-xs bg-base-200 border border-base-300 p-4 rounded-box">
      <legend className="fieldset-legend">Login</legend>
        <form className="flex flex-col gap-2" onSubmit={handleSubmit}>
          <label className="fieldset-label">Email</label>
          <input
            type="email"
            className="input validator" 
            placeholder="Email"
            required
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
          <div className="validator-hint hidden">Enter valid email address</div>
      
          <label className="fieldset-label">Password</label>
          <input
            type="password"
            className="input"
            required
            placeholder="Password"
            pattern="(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,}"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
      
          <button
            type="submit"
            className={`btn btn-neutral mt-4 ${loading ? 'loading' : ''}`}
            disabled={loading}
          >
            {loading ? 'Logging in...' : 'Login'}
          </button>
      </form>
    </fieldset>
  )
}

export default Login
