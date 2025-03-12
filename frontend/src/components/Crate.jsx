import { buyCrate } from "../api/store"
import useInventory from "../stores/inventoryStore"

function Crate({ name, count }) {
  const { inventory, setInventory, addItem, removeItem } = useInventory()

  const handleSubmit = async () => {
    try {
      const jwt = localStorage.getItem("jwt")
      const data = await buyCrate(jwt, name, count)
      data.forEach(item => {
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
          {name} {count}
        </h2>
        <figure>
          <img src="https://img.daisyui.com/images/stock/photo-1606107557195-0e29a4b5b4aa.webp" alt="" />
        </figure>
        <div className="card-actions">
          <button className="btn btn-primary" onClick={handleSubmit}>Buy crate</button>
        </div>
      </div>
    </div>
  )
}

export default Crate
