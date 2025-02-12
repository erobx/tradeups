

function Footer() {
  return (
    <footer className="footer sm:footer-horizontal bg-neutral text-neutral-content p-4 mt-auto w-full">
      <aside className="grid-flow-col items-center">
        <p>Copyright @ {new Date().getFullYear()} - All right reserved</p>
      </aside>
    </footer>
  )
}

export default Footer
