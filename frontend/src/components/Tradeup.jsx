import { useEffect } from "react"
import { useParams } from "react-router"
import TradeupGrid from "./TradeupGrid"

function Tradeup() {
  const params = useParams()
  const tradeupId = params.tradeupId
  const status = "Waiting"

  const fetchTradeup = async () => {

  }

  useEffect(() => {
    console.log(tradeupId)
  }, [])

  return (
    <div className="flex flex-col items-center mt-5">
      <div className="font-bold text-xl">Trade Up Status:&nbsp;
        <span className="text-info">{status}</span>
      </div>
      <TradeupGrid />
    </div>
  )
}

export default Tradeup
