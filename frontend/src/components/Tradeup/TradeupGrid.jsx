import { useMemo } from "react"
import StatTrakBadge from "../StatTrakBadge"
import TradeupModal from "./TradeupModal"
import useUser from "../../stores/userStore"
import { removeSkinFromTradeup } from "../../api/tradeups"
import useInventory from "../../stores/inventoryStore"

function Modal({ invId, tradeupId }) {
  const { inventory, setInventory, addItem, removeItem } = useInventory()

  const onClick = async () => {
    const jwt = localStorage.getItem("jwt")
    const data = await removeSkinFromTradeup(jwt, invId, tradeupId)
    if (data) {
      addItem(data)
    } else {
      return
    }
    console.log(`Removed skin ${invId} from tradeup ${tradeupId}`)

    // add item back to inventory
    //addItem()
  }

  return (
    <dialog id={`modal_${invId}`} className="modal">
      <div className="modal-box max-h-3xl">
        <h3 className="font-bold text-lg mb-2">Remove skin from Tradeup?</h3>
        <button className="btn btn-error" onClick={onClick}>Remove skin</button>
      </div>
      <form method="dialog" className="modal-backdrop">
        <button>close</button>
      </form>
    </dialog>
  )
}

function GridItem({ id, tradeupId, name, wear, price, isStatTrak, imgSrc, owned }) {
  const onSelect = () => {
    if (owned) {
      document.getElementById(`modal_${id}`).showModal()
    }
  }

  return (
    <div
      className="card card-xs w-48 bg-base-200 shadow-md m-0.5 hover:outline-4 outline-info"
      onClick={onSelect}
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
      <Modal invId={id} tradeupId={tradeupId} />
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

  const skinsWithOwnership = useMemo(() => {
    return skins.map(skin => ({
      ...skin,
      owned: skin.userId == user.id
    }))
  }, [skins, user.id])

  return (
    <div className="grid grid-cols-5 grid-rows-2 rounded mt-5 gap-2">
      {skinsWithOwnership.map(s => (
        <GridItem
          key={s.inventoryId}
          id={s.inventoryId}
          tradeupId={tradeupId}
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
