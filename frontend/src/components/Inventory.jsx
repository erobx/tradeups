import { useState } from "react"
import InventoryItem from "./InventoryItem"
import EmptyItem from "./EmptyItem"
import useInventory from "../stores/inventoryStore"

function Inventory() {
  const { inventory, setInventory, addItem, removeItem } = useInventory()
  const [loading, setLoading] = useState(false)
  
  if (loading) {
    return (
      <div className="flex justify-center items-center min-h-screen">
        <div className="loading loading-spinner loading-xl"></div>
      </div>
    )
  }

  if (inventory.length == 0) {
    return (
      <div>
        <EmptyItem />
      </div>
    )
  }

  return (
    <div className="grid grid-flow-row lg:grid-cols-7 gap-4 md:grid-cols-2">
      {inventory.map((item, index) => (
        <div key={index} className="item">
          <InventoryItem 
            id={item.id}
            name={item.name}
            rarity={item.rarity}
            wear={item.wear}
            price={item.skinPrice}
            isStatTrak={item.isStatTrak}
            imgSrc={item.imageSrc}
          />
        </div>
      ))}
    </div>
  )
}

export default Inventory
