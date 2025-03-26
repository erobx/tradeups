import useUser from "../stores/userStore"
import useInventory from "../stores/inventoryStore"

function Logout({ setLoggedIn }) {
  const { user, setUser, setBalance } = useUser()
  const { inventory, setInventory, addItem, removeItem } = useInventory()

  // TODO: change for cookies when that finally works
  const handleLogout = async () => {
    localStorage.removeItem("jwt")
    setLoggedIn(false)
    setUser({})
    setInventory([])
  }

  return (
      <span onClick={handleLogout}>Logout</span>
  )
}

export default Logout
