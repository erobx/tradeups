import { useEffect, useState } from "react"
import { textMap } from "../../constants/rarity"
import { useNavigate } from "react-router"
import { getRecentTradeups } from "../../api/user"

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
      default:
        setStatusColor("text-accent")
    }
  }, [status])

  return (
    <li className="list-row">
      <div>
        <div className="font-bold">Rarity</div>
        <div className={`${textColor} font-bold`}>{rarity}</div>
      </div>

      <div>
        <div className="font-bold">Status</div>
        <div className={`${statusColor} font-bold`}>{status}</div>
      </div>

      <div className="min-w-[300px] max-h-[100px] overflow-y-auto">
        <div className="font-bold sticky top-0 bg-base-300 z-10 pb-2">Skins Invested</div>
        {skins.map((s, index) => (
          <div
            key={s.inventoryId}
            className={`
              ${index % 2 === 0 ? 'text-accent' : ''}
            `}>
            {s.name} ({s.wear})
          </div>
        ))}
      </div>

      <div className="ml-10 mr-10">
        <div className="font-bold">Value</div>
        <div className="font-bold text-primary">${value.toFixed(2)}</div>
      </div>

      <div className="mr-8">
        <button className="btn btn-soft btn-warning" onClick={logTradeupId}>View</button>
      </div>
    </li>
  )
}

function RecentTradeups({ user }) {
  const [rows, setRows] = useState([])
  const [isLoading, setIsLoading] = useState(true)
  const [error, setError] = useState(null)

  const fetchRecentTradeups = async () => {
    try {
      setIsLoading(true)
      const jwt = localStorage.getItem("jwt")
      const data = await getRecentTradeups(jwt, user.id)
      setRows(data)
    } catch (error) {
      console.error("Failed to fetch recent tradeups: ", error);
      setError(error.message)
    } finally {
      setIsLoading(false)
    }
  }

  useEffect(() => {
    fetchRecentTradeups()
  }, [])

  if (isLoading) {
    return (
      <div className="w-fit bg-base-300 rounded-box shadow-md">
        <div className="p-4 text-center">Loading trade ups...</div>
      </div>
    )
  }

  if (error) {
    return (
      <div className="w-fit bg-base-300 rounded-box shadow-md">
        <div className="p-4 text-red-500">Error: {error}</div>
      </div>
    )
  }

  if (rows.length === 0) {
    return (
      <div className="w-fit bg-base-300 rounded-box shadow-md">
        <div className="p-4 text-center">No recent trade ups found</div>
      </div>
    )
  }

  return (
    <div className="w-fit bg-base-300 rounded-box shadow-md">
      <ul className="list">
        <div className="flex items-center">
          <li className="p-4 pb-2 text-sm opacity-70 tracking-wide">Recent Trade Ups</li>
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
