import { useRef } from "react";
import { useDraggable } from "src/hooks/useDraggable";

import styles from "src/styles/browser/index.module.css";
import {
  ClassicGameCard,
  MegaDiamondGameCard,
  MexicanChilliGameCard,
} from "./game";

const Browser = () => {
  const ref = useRef<HTMLElement>() as React.MutableRefObject<HTMLElement>;
  const { events } = useDraggable(ref, {
    decayRate: 0.97,
    safeDisplacement: 100,
    applyRubberBandEffect: true,
    activeMouseButton: "Right",
  });

  return (
    <div ref={ref as any} {...events} className={styles.gameBrowser}>
      <ClassicGameCard />
      <div className={styles.gameDivider} />
      <MegaDiamondGameCard />
      <MexicanChilliGameCard />
    </div>
  );
};

export default Browser;
