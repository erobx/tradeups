import { Routes, Route } from "react-router"

import DashboardDrawer from "../components/Dashboard/DashboardDrawer"
import Inventory from "../components/Inventory/Inventory"
import Dashboard from "../components/Dashboard/Dashboard"

function DashboardPage() {
  return (
    <div className="flex justify-evenly gap-2">
      <div>
        <DashboardDrawer />
      </div>
      <div className="mt-2">
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/inventory" element={<Inventory />} />
        </Routes>
      </div>
    </div>
  )
}

export default DashboardPage
