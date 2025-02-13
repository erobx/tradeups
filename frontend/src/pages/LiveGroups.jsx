import { useState } from "react"
import TradeupRow from "../components/TradeupRow"

function LiveGroups() {
  const liveTradeups = useState([])
  const tempT = [
    {id: 0},
    {id: 1},
    {id: 2},
    {id: 3},
  ]

  return (
    <div className="flex flex-col items-center gap-2 mt-3">
      <h1 className="text-warning text-3xl font-bold">Live Groups</h1>
      <div>Filter placeholder</div>
      {tempT.map(t => <TradeupRow id={t.id} />)}
    </div>
  )
}

export default LiveGroups
