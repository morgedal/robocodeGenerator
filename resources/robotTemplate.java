package sample;

import robocode.HitByBulletEvent;
import robocode.Robot;
import robocode.ScannedRobotEvent;
import robocode.HitWallEvent;
import robocode.HitRobotEvent;
import java.util.Random;

public class ROBOT_NAME extends Robot {

	public void run() {

		while (true) {
			BODY
		}
	}

	public void onScannedRobot(ScannedRobotEvent e) {
		SCANNED
	}

	public void onHitByBullet(HitByBulletEvent e) {
		ON_BULLET_HIT
	}

	public void onHitWall(HitWallEvent e) {
		ON_HIT_WALL
	}

	public void onHitRobot(HitRobotEvent e) {
		ON_HIT_ROBOT
	}

	public int random(int low, int high) {
		Random r = new Random();
		return r.nextInt(high-low) + low;
	}
}												

