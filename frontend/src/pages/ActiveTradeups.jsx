import { useEffect } from "react"
import TradeupRow from "../components/TradeupRow"
import useTradeups from "../stores/tradeupsStore"

function ActiveTradeups() {
  const { tradeups, setTradeups, addTradeup, removeTradeup } = useTradeups()

  return (
    <div className="flex flex-col items-center gap-2 mt-3">
      <h1 className="text-warning text-3xl font-bold">Active Trade Ups</h1>
      <div>Filter placeholder</div>
      {tradeups.map((t, index) => (
        <TradeupRow
          key={index}
          id={t.id}
          players={t.players}
          rarity={t.rarity}
          skins={t.skins}
          status={t.status}
        />
      ))}
    </div>
  )
}

export default ActiveTradeups
