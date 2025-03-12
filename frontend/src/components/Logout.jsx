import { useState } from "react"


function Logout({ setLoggedIn }) {
  const [loading, setLoading] = useState(false)

  // TODO: change for cookies when that finally works
  const handleLogout = async () => {
    setLoading(true)
    localStorage.removeItem("jwt")
    localStorage.removeItem("userId")
    setLoggedIn(false)
    setLoading(false)
  }

  return (
      <span onClick={handleLogout}>Logout</span>
  )
}

export default Logout
