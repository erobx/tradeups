import StatTrakBadge from "../StatTrakBadge"
import { addSkinToTradeup } from "../../api/tradeups"

function ModalItem({ invId, tradeupId, name, wear, isStatTrak, imgSrc, price, removeItem }) {

  const addSkin = async () => {
    console.log(`adding skin ${invId} to tradeup ${tradeupId}...`)
    const jwt = localStorage.getItem("jwt")
    try {
      const res = await addSkinToTradeup(jwt, invId, tradeupId)
      if (res.status !== 201) {
        return
      }
      removeItem(invId)
    } catch (error) {
      console.error("Error: ", error)
    }
  }

  return (
    <div className={`card card-md w-56 h-48 bg-base-300 hover:border-4 hover:cursor-pointer`} onClick={addSkin}>
      <h1 className="text-sm font-bold text-primary ml-1.5 mt-0.5">${price}</h1>
      <figure>
        <img
          alt={imgSrc}
          src={imgSrc}
          width={100}
          height={50}
        />
      </figure>
      <div className="card-body items-center">
        <h1 className="card-title text-sm">{name}</h1>
        <h1 className="card-title text-xs">({wear})</h1>
        <div className="flex gap-2">
          {isStatTrak && <StatTrakBadge />}
        </div>
      </div>
    </div>
  )
}

export default ModalItem
