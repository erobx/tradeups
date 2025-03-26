import { buyCrate } from "../api/store"
import useInventory from "../stores/inventoryStore"

function Crate({ name, count }) {
  const { inventory, setInventory, addItem, removeItem } = useInventory()

  const handleSubmit = async () => {
    try {
      const jwt = localStorage.getItem("jwt")
      const data = await buyCrate(jwt, name, count)
      data.forEach(item => {
        // doesn't work
        item.price = Number(item.price.toFixed(2))
        addItem(item)
      })
    } catch (error) {
      console.error(error)
    }
  }

  return (
    <div className="card bg-base-300 w-52 shadow-sm">
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
