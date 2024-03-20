# runs dmenu based on yml structure

build this and put into your bin folder, also works with rofi if you comment line 30 and uncomment line 31

use:


dmylm test.yml


```yml
---
option: command
option2: command
submenu:
  option3: command
  submenu:
    option3: command
  submenu2:
    option4: command
    option5: command
```


