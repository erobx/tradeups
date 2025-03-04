import InventoryItem from "./InventoryItem"
import EmptyItem from "./EmptyItem"
import useInventory from "../stores/inventoryStore"
import { usePresignedUrls } from "../hooks/usePresignedUrls"

function Inventory() {
  const { inventory, setInventory, addItem, removeItem } = useInventory()
  const processedInventory = usePresignedUrls(inventory)
  
  if (processedInventory.length == 0) {
    return (
      <div>
        <EmptyItem />
      </div>
    )
  }

  return (
    <div className="grid grid-flow-row lg:grid-cols-7 gap-4 md:grid-cols-2">
      {processedInventory.map((item, index) => (
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
