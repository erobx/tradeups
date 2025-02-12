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
            Skins in Pool:
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
    <div className="card card-sm">
      <div className="card-body">
        <div className="card-title">
          <button className="btn btn-soft rounded-md btn-success" onClick={onJoin}>
            Join
          </button>
        </div>
      </div>
    </div>
  )
}

function TradeupRow({ id }) {
  const [rarity, setRarity] = useState("Covert")
  const [countItems, setCountItems] = useState(10)
  const [totalPrice, setTotalPrice] = useState(1.69)
  const [players, setPlayers] = useState([])

  const tempPlayers = [
    {id: 0, initials: "ER", imgSrc: ""},
    {id: 0, initials: "JD", imgSrc: ""}
  ]

  const dividerColor = dividerMap[rarity]

  const getTradeup = async () => {
    
  }

  useEffect(() => {
    // getTradeup()
  },[])

  return (
    <div className="join bg-base-300 lg:w-3/4 rounded-md">
      <InfoPanel className="join-item" count={countItems} />
      <div className={`divider divider-horizontal ${dividerColor}`}></div>

      <ItemCarousel className="join-item"/>
      <div className="divider divider-horizontal divider-secondary"></div>

      <DetailsPanel className="join-item" total={totalPrice} />
      <div className="divider divider-horizontal divider-warning"></div>

      <PlayersPanel className="join-item" players={tempPlayers} />
      <div className="divider divider-horizontal divider-primary"></div>

      <ButtonPanel
        tradeupId={id}
      />
    </div>
  )
}

export default TradeupRow
