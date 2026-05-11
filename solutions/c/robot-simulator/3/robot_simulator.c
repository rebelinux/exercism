#include "robot_simulator.h"
#include "string.h"

robot_status_t robot_create(robot_direction_t direction, int x, int y) {
    robot_position_t rp = {
        .x = x,
        .y = y
    };
    robot_status_t rs = {
        .direction = direction,
        .position = rp
    };
    return rs;
}

void robot_move(robot_status_t *robot, const char *commands) {
    int l = strlen(commands);
    for (int i = 0; i < l; i++) {
        switch (commands[i])
        {
            case 'R':
                if (robot->direction == DIRECTION_NORTH) {
                    robot->direction = (robot->direction + 1) % DIRECTION_MAX;
                } else if (robot->direction == DIRECTION_SOUTH) {
                    robot->direction = (robot->direction + 1) % DIRECTION_MAX;
                } else if (robot->direction == DIRECTION_WEST) {
                    robot->direction = (robot->direction + 1) % DIRECTION_MAX;
                } else if (robot->direction == DIRECTION_EAST) {
                    robot->direction = (robot->direction + 1) % DIRECTION_MAX;
                }
                break;
            case 'L':
                if (robot->direction == DIRECTION_NORTH) {
                    robot->direction = (robot->direction - 1) % DIRECTION_MAX;
                } else if (robot->direction == DIRECTION_SOUTH) {
                    robot->direction = (robot->direction - 1) % DIRECTION_MAX;
                } else if (robot->direction == DIRECTION_WEST) {
                    robot->direction = (robot->direction - 1) % DIRECTION_MAX;
                } else if (robot->direction == DIRECTION_EAST) {
                    robot->direction = (robot->direction - 1) % DIRECTION_MAX;
                }
                break;
            case 'A':
                if (robot->direction == DIRECTION_NORTH) {
                    robot->position.y++;
                } else if (robot->direction == DIRECTION_SOUTH) {
                    robot->position.y--;
                } else if (robot->direction == DIRECTION_EAST) {
                    robot->position.x++;
                } else if (robot->direction == DIRECTION_WEST) {
                    robot->position.x--;
                }
                break;
        
            default:
                break;
        }
    }
}