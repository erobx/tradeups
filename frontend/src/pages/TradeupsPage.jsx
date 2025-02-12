import { useState } from "react"
import TradeupRow from "../components/TradeupRow"

function TradeupsPage() {
  const activeTradeups = useState([])

  return (
    <div className="flex flex-col items-center gap-1.5 mt-3">
      <h1 className="text-warning text-3xl font-bold">Active Groups</h1>
      <div>Filter placeholder</div>
      <TradeupRow
        id={0}
      />
      <TradeupRow
        id={1}
      />
    </div>
  )
}

export default TradeupsPage
