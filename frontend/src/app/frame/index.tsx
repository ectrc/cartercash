import { useEffect } from "react";
import { Outlet } from "@tanstack/react-router";

import styles from "src/styles/frame/index.module.css";
import FrameMoney from "./money";
import FrameLevel from "./level";

const Frame = () => {
  const noop = (e: MouseEvent) => e.preventDefault();
  useEffect(() => {
    document.addEventListener("contextmenu", noop);
    return () => {
      document.removeEventListener("contextmenu", noop);
    };
  }, []);

  return (
    <div className={styles.frame}>
      <div className={styles.frameMenu}>
        <FrameMoney />
        <FrameLevel />
      </div>
      <section>
        <Outlet />
      </section>
    </div>
  );
};

export default Frame;
