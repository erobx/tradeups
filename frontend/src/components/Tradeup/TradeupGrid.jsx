import { useEffect } from "react"
import StatTrakBadge from "../StatTrakBadge"
import TradeupModal from "./TradeupModal"
import useInventory from "../../stores/inventoryStore"
import useUser from "../../stores/userStore"

function GridItem({ name, wear, price, isStatTrak, imgSrc, owned }) {
  const onClick = async () => {
    console.log("User owns: ", owned)
  }

  return (
    <div
      className="card card-xs w-48 bg-base-200 shadow-md m-0.5 hover:outline-4 outline-info"
      onClick={onClick}
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

function EmptyGridItem({ tradeupId, rarity }) {
  return (
    <div className="card card-xs w-48 bg-base-200 shadow-md m-0.5">
      <div className="card-body items-center">
        <div className="card-actions mt-12">
          <TradeupModal tradeupId={tradeupId} rarity={rarity} />
        </div>
      </div>
    </div>
  )
}

function TradeupGrid({ tradeupId, rarity, skins }) {
  const { user, setUser } = useUser()

  useEffect(() => {
    skins = skins.map(skin => ({
      ...skin,
      owned: skin.userId === user.id
    }))
  }, [])

  return (
    <div className="grid grid-cols-5 grid-rows-2 rounded mt-5 gap-2">
      {skins.map(s => (
        <GridItem
          key={s.inventoryId}
          name={s.name}
          wear={s.wear}
          price={s.price}
          isStatTrak={s.isStatTrak}
          imgSrc={s.imageSrc}
          owned={s.owned}
        />
      ))}
      {skins.length < 10 && (
        Array.from({ length: 10 - skins.length }).map((_, index) => (
          <EmptyGridItem key={`empty-${index}`} tradeupId={tradeupId} rarity={rarity} />
      )))}
    </div>
  )
}

export default TradeupGrid
