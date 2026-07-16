public class CarsAssemble {

    public double productionRatePerHour(int speed) {
        if (speed == 0) {
            return 0;
        } else if (speed >= 1 && speed <= 4) {
            return speed * 221;
        } else if (speed >= 5 && speed <= 8) {
            return speed * 221 * 0.9;
        } else if (speed == 9) {
            return speed * 221 * 0.8;
        } else if (speed == 10) {
            return speed * 221 * 0.77;
        } else {
            throw new Error("Input a valid number");
        }
    }

    public int workingItemsPerMinute(int speed) {
        return (int)(productionRatePerHour(speed) / 60);
    }
}
