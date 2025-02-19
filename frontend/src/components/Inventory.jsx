import { useEffect, useState } from "react"
import InventoryItem from "./InventoryItem"
import EmptyItem from "./EmptyItem"
import useUserId from "../stores/userStore"
import { getInventory } from "../api/inventory"

function Inventory() {
  const { userId, setUserId } = useUserId()
  const [items, setItems] = useState([])
  const [loading, setLoading] = useState(false)

  const loadItems = async (userId) => {
    setLoading(true)
    const jwt = localStorage.getItem("jwt")
    try {
      const data = await getInventory(jwt, userId)
      setItems(data.skins)
    } catch (error) {
      console.error(error)
    } finally {
      setLoading(false)
    }
  }
  
  useEffect(() => {
    loadItems(userId)
  }, [userId])

  if (loading) {
    return (
      <div className="flex justify-center items-center min-h-screen">
        <div className="loading loading-spinner loading-xl"></div>
      </div>
    )
  }

  if (items.length == 0) {
    return (
      <div>
        <EmptyItem />
      </div>
    )
  }

  return (
    <div className="grid grid-flow-row lg:grid-cols-7 gap-4 md:grid-cols-2">
      {items.map((item, index) => (
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
