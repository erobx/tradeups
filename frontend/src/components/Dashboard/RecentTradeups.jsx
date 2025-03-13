import { useEffect, useState } from "react"
import { textMap } from "../../constants/rarity"
import { useNavigate } from "react-router"

function ListRow({ tradeupId, rarity, status, skins, value }) {
  const navigate = useNavigate()
  const textColor = textMap[rarity]
  const [statusColor, setStatusColor] = useState("")

  const logTradeupId = () => {
    navigate(`/tradeups/${tradeupId}`)
  }

  useEffect(() => {
    switch (status) {
      case "Active":
        setStatusColor("text-info")
        break
      case "Completed":
        setStatusColor("text-success")
        break
    }
  }, [status])

  return (
    <li className="list-row">
      <div>
        <div className="font-bold">Rarity</div>
        <div className={`${textColor}`}>{rarity}</div>
      </div>

      <div>
        <div className="font-bold">Status</div>
        <div className={`${statusColor}`}>{status}</div>
      </div>

      <div className="min-w-[300px]">
        <div className="font-bold">Skins Invested</div>
        {skins.map(s => (
          <div key={s.id} className="text-warning">{s.name} ({s.wear})</div>
        ))}
      </div>

    <div>
        <div className="font-bold">Value</div>
        <div className="text-primary">${value}</div>
      </div>

      <div>
        <button className="btn btn-soft btn-primary" onClick={logTradeupId}>Go To</button>
      </div>
    </li>
  )
}

function RecentTradeups() {
  const skins1 = [
    {id: 0, name: "AK-47 | Slate", wear: "Factory New"},
    {id: 1, name: "Desert Eagle | Printstream", wear: "Factory New"},
  ]

  const skins2 = [
    {id: 0, name: "M4A4 | Howl", wear: "Factory New"},
    {id: 1, name: "M4A1-S | Black Lotus", wear: "Minimal Wear"},
  ]

  const rows = [
    {id: 1, rarity: "Consumer", status: "Active", skins: skins1, value: 35.21},
    {id: 2, rarity: "Industrial", status: "Completed", skins: skins2, value: 121.01},
  ]

  return (
    <div className="w-fit bg-base-300 rounded-box shadow-md">
      <ul className="list">
        <div className="flex items-center">
          <li className="p-4 pb-2 text-sm opacity-70 tracking-wide">Recent Trade Ups</li>
          {/* TODO: filter by rarity, status, money
          <li>
          </li>
          */}
        </div>
        {rows.map((r, index) => (
          <ListRow
            key={index}
            tradeupId={r.id}
            rarity={r.rarity}
            status={r.status}
            skins={r.skins}
            value={r.value}
          />
        ))}
      </ul>
    </div>
  )
}

export default RecentTradeups
