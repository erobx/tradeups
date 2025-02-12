import RecentTradeups from "./RecentTradeups"
import Stats from "./Stats"

function Dashboard() {
  return (
    <div className="flex gap-10">
      <div>
        <RecentTradeups />
      </div>
      <div>
        <Stats />
      </div>
    </div>
  )
}

export default Dashboard
