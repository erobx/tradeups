import { buyCrate } from "../api/store"
import useInventory from "../stores/inventoryStore"
import useUser from "../stores/userStore"

function Crate({ name, rarity, count }) {
  const { inventory, setInventory, addItem, removeItem } = useInventory()
  const { user, setUser, setBalance } = useUser()

  const handleSubmit = async () => {
    try {
      const jwt = localStorage.getItem("jwt")
      const data = await buyCrate(jwt, name, rarity, count)
      // data {skins: [], balance: 0.00}
      data.skins.forEach(item => {
        addItem(item)
      })
      setBalance(data.balance)
    } catch (error) {
      console.error(error)
    }
  }

  return (
    <div className="card bg-base-300 w-64 shadow-sm">
      <div className="card-body">
        <h2 className="card-title">
          {name}
        </h2>
        <figure>
          <img src="crate.png" alt="" />
        </figure>
        <div className="card-actions">
          <button className="btn btn-primary" onClick={handleSubmit}>Buy crate</button>
        </div>
      </div>
    </div>
  )
}

export default Crate
