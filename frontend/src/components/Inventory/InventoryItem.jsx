import RarityBadge from "../RarityBadge"
import StatTrakBadge from "../StatTrakBadge"
import { outlineMap } from "../../constants/rarity"

function InventoryItem({ name, rarity, wear, price, isStatTrak, imgSrc }) {
  const outlineColor = outlineMap[rarity]

  return (
    <div
      className={`card card-xs w-52 bg-base-200 shadow-md hover:outline-4 ${outlineColor}`}
    >
      <h1 className="text-primary ml-1.5">${price}</h1>
      <figure>
        <img
          loading="eager"
          alt={name}
          src={imgSrc}
        />
      </figure>
      <div className="card-body items-center">
        <h1 className="card-title text-xs">{name}</h1>
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
