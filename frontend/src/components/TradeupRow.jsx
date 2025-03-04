import { useEffect, useState } from "react"
import { useNavigate } from "react-router"
import { dividerMap } from "../constants/rarity"
import ItemCarousel from "./ItemCarousel"
import AvatarGroup from "./AvatarGroup"

function InfoPanel({ count }) {
  return (
    <div className="card card-sm">
      <div className="card-body justify-center ml-4">
        <div className="flex flex-col items-center">
          <div className="card-title font-bold text-xl">
            Skin Count:
          </div>
          <div className="card-title font-bold text-accent text-xl">
            {count}
          </div>
        </div>
      </div>
    </div>
  )
}

function DetailsPanel({ total }) {
  return (
    <div className="card card-sm">
      <div className="card-body justify-center">
        <div className="flex flex-col items-center">
          <div className="card-title">
            Pool Value:
          </div>
          <div className="card-title">
            ${total}
          </div>
        </div>
      </div>
    </div>
  )
}

function PlayersPanel({ players }) {
  return (
    <div className="card card-sm">
      <div className="card-body justify-center">
        <div className="flex flex-col items-center">
          <div className="card-title">
            Players:
          </div>
          <div className="card-title">
            <AvatarGroup
              avatars={players}
            />
          </div>
        </div>
      </div>
    </div>
  )
}

function ButtonPanel({ tradeupId }) {
  const navigate = useNavigate()

  const onJoin = () => {
    const url = `/tradeups/${tradeupId}`
    navigate(url)
  }

  return (
    <button className="btn btn-soft rounded-md btn-success w-40" onClick={onJoin}>
      Join
    </button>
  )
}

function TradeupRow({ id, players, rarity, skins, status }) {
  const [totalPrice, setTotalPrice] = useState(0)

  const tempPlayers = [
    {id: 0, initials: "ER", imgSrc: ""},
    {id: 1, initials: "JD", imgSrc: ""}
  ]

  const dividerColor = dividerMap[rarity]

  useEffect(() => {
    let sum = totalPrice
    skins.forEach(skin => {
      sum += parseFloat(skin.price)
    })
    setTotalPrice(sum)
  }, [])

  return (
    <div className="join bg-base-300 border-6 border-base-200 items-center lg:w-3/4 rounded-md">
      <InfoPanel className="join-item" count={skins.length} />
      <div className={`divider divider-horizontal ${dividerColor}`}></div>

      <ItemCarousel className="join-item" skins={skins} />
      <div className="divider divider-horizontal divider-secondary"></div>

      <DetailsPanel className="join-item" total={totalPrice} />
      <div className="divider divider-horizontal divider-warning"></div>

      <PlayersPanel className="join-item" players={tempPlayers} />
      <div className="divider divider-horizontal divider-primary"></div>

      <ButtonPanel className="join-item" tradeupId={id} />
    </div>
  )
}

export default TradeupRow
