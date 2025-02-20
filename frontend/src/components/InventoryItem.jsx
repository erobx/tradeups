import RarityBadge from "./RarityBadge"
import StatTrakBadge from "./StatTrakBadge"
import { borderMap } from "../constants/rarity"

function InventoryItem({ id, name, rarity, wear, price, isStatTrak, imgSrc }) {
  const borderColor = borderMap[rarity]

  const testClick = () => {
    console.log(id)
  }

  return (
    <div
      className={`card card-xs w-48 bg-base-200 shadow-md hover:border-r-4 ${borderColor}`}
      onClick={testClick}
    >
      <h1 className="text-primary ml-1.5">${price}</h1>
      <figure>
        <img
          alt={name}
          src={imgSrc}
        />
      </figure>
      <div className="card-body items-center">
        <h1 className="card-title text-sm">{name}</h1>
        <h2 className="card-title text-xs">({wear})</h2>
        <div className="flex gap-2">
          <div>
            <RarityBadge
              rarity={rarity}
            />
          </div>
          {isStatTrak && <StatTrakBadge />}
        </div>
      </div>
    </div>
  )
}

export default InventoryItem
