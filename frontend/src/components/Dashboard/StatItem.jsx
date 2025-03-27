import { LazyLoadImage } from "react-lazy-load-image-component"
import RarityBadge from "../RarityBadge"
import StatTrakBadge from "../StatTrakBadge"

function StatItem({ name, wear, rarity, isStatTrak, imgSrc, price }) {
  return (
    <div className="card card-xs w-48">
      <h1 className="ml-1.5">${price}</h1>
      <figure>
        <div>
          <LazyLoadImage
            alt={imgSrc}
            src={imgSrc}
            width={100}
            height={50}
          />
        </div>
      </figure>
      <div className="card-body items-center">
        <h1 className="card-title text-sm">{name}</h1>
        <h1 className="card-title text-xs">({wear})</h1>
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

export default StatItem
