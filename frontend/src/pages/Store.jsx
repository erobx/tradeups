import Crate from "../components/Crate"

function Store() {
  return (
    <div className="flex flex-col gap-2 mt-5">
      <div className="flex flex-col items-center text-center">
        <h1 className="font-bold text-3xl">Welcome to the Store!</h1>
        <p>
          Here you can buy crates of Consumer, Industrial, and Mil-Spec quality
          <br/>skins to use in Trade Ups.
        </p>
      </div>
      <div className="flex justify-start ml-4">
        <div className="grid grid-flow-row grid-cols-5 gap-6">
          <Crate name="Consumer" count={3} />
          <Crate name="Industrial" count={3} />
          <Crate name="Mil-Spec" count={3} />
        </div>
      </div>
    </div>
  )
}

export default Store
