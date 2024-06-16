import styles from "src/styles/frame/level.module.css";

const FrameLevel = () => {
  return (
    <div className={styles.frameLevelContainer}>
      <section className={styles.frameAvatar}></section>
      <div className={styles.frameLevel}>
        <span>231</span>
        <div></div>
      </div>
    </div>
  );
};

export default FrameLevel;
