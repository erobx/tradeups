import { useEffect } from "react"
import { useParams } from "react-router"
import TradeupGrid from "./TradeupGrid"

function Tradeup() {
  const params = useParams()
  const tradeupId = params.tradeupId
  const status = "Active"
  const [rarity, setRarity] = "Consumer"

  const fetchTradeup = async () => {

  }

  useEffect(() => {
  }, [])

  return (
    <div className="flex flex-col items-center mt-5">
      <div className="font-bold text-xl">Trade Up Status:&nbsp;
        <span className="text-info">{status}</span>
      </div>
      <TradeupGrid rarity={rarity} />
    </div>
  )
}

export default Tradeup
