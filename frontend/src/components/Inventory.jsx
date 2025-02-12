import { useEffect, useState } from "react"
import InventoryItem from "./InventoryItem"
import EmptyItem from "./EmptyItem"

function Inventory() {
  const [items, setItems] = useState([])

  const loadItems = async () => {

  }
  
  useEffect(() => {
    //loadItems()
  }, [])

  const tempItems = [
    {id: 0, name: "M4A4 | Howl", rarity: "Contraband", isStatTrak: true, imgSrc: "/m4a4-howl.png"},
    {id: 1, name: "Desert Eagle | Printstream", rarity: "Covert", isStatTrak: true, imgSrc: "/de-printstream.png"},
    {id: 2, name: "M4A1-S | Black Lotus", rarity: "Classified", isStatTrak: false, imgSrc: "/m4a1-black-lotus.png"},
    {id: 3, name: "AK-47 | Slate", rarity: "Restricted", isStatTrak: false, imgSrc: "/ak-slate.png"},
    {id: 4, name: "AUG | Wings", rarity: "Mil-Spec", isStatTrak: false, imgSrc: "/aug-wings.png"},
    {id: 5, name: "Five-SeveN | Candy Apple", rarity: "Industrial", isStatTrak: false, imgSrc: "/fs-candy-apple.png"},
    {id: 6, name: "SG 553 | Tornado", rarity: "Consumer", isStatTrak: false, imgSrc: "/sg-tornado.png"},
  ]

  if (tempItems.length == 0) {
    return (
      <div>
        <EmptyItem />
      </div>
    )
  }

  return (
    <div className="grid grid-flow-row lg:grid-cols-7 gap-4 md:grid-cols-2">
      {tempItems.map((item, index) => (
        <div key={index} className="item">
          <InventoryItem 
            id={item.id}
            name={item.name}
            rarity={item.rarity}
            isStatTrak={item.isStatTrak}
            imgSrc={item.imgSrc}
          />
        </div>
      ))}
    </div>
  )
}

export default Inventory
