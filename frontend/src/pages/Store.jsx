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

      <div className="divider"></div>

      <div className="flex flex-col items-center gap-2">

        <h1 className="font-bold text-xl ml-4.5">Crates containing 3 skins:</h1>
        <div className="flex justify-start ml-4">
          <div className="flex gap-6">
            <Crate name="Consumer" count={3} />
            <Crate name="Industrial" count={3} />
            <Crate name="Mil-Spec" count={3} />
          </div>
        </div>

        <div className="divider"></div>

        <h1 className="font-bold text-xl ml-4.5">Crates containing 5 skins:</h1>
        <div className="flex justify-start ml-4">
          <div className="flex gap-6">
            <Crate name="Consumer" count={5} />
            <Crate name="Industrial" count={5} />
            <Crate name="Mil-Spec" count={5} />
          </div>
        </div>

      </div>
    </div>
  )
}

export default Store
