import StatTrakBadge from "./StatTrakBadge"
import InventoryModal from "./InventoryModal"

function GridItem({ name, wear, price, isStatTrak, imgSrc }) {
  const testClick = () => {

  }

  return (
    <div
      className="card card-xs w-48 bg-base-200 shadow-md m-0.5"
      onClick={testClick}
    >
      <h1 className="ml-1.5">${price}</h1>
      <figure>
        <img
          alt={name}
          src={imgSrc}
        />
      </figure>
      <div className="card-body items-center">
        <h1 className="card-title text-sm">{name}</h1>
        <h1 className="card-title text-sm">({wear})</h1>
        <div className="flex gap-2">
          {isStatTrak && <StatTrakBadge />}
        </div>
      </div>
    </div>
  )
}

function EmptyGridItem({ rarity }) {
  return (
    <div className="card card-xs w-48 bg-base-200 shadow-md m-0.5">
      <div className="card-body items-center">
        <div className="card-actions mt-12">
          <InventoryModal rarity={rarity} />
        </div>
      </div>
    </div>
  )
}

function TradeupGrid({ rarity }) {
  const skins = [
    {id: 0, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: true, imgSrc: "/aug-wings.png"},
    {id: 1, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
    {id: 2, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: true, imgSrc: "/aug-wings.png"},
    {id: 3, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
    {id: 4, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
  ]

  return (
    <div className="grid grid-cols-5 grid-rows-2 rounded mt-5 gap-2">
      {skins.map((s, index) => (
        <GridItem
          key={index}
          name={s.name}
          wear={s.wear}
          price={s.price}
          isStatTrak={s.isStatTrak}
          imgSrc={s.imgSrc}
        />
      ))}
      {skins.length < 10 && (
        Array.from({ length: 10 - skins.length }).map((_, index) => (
          <EmptyGridItem key={`empty-${index}`} rarity={rarity} />
      )))}
    </div>
  )
}

export default TradeupGrid
